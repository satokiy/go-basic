/*
 mainゴルーチンで, task関数の実行を監視する
 task関数はheartbeatをmain関数側に送り続け、一定時間送られないとmain関数はタイムアウトと判断する
*/

/*
 プログラムの作成順序
 1. ログ・ファイルの作成
   - main関数実行の最初にログファイルを作り、main関数終了時にログファイルを閉じる
 2. task関数の作成
   - task関数はmain関数にheartbeatを送り続ける
 3. main関数の作成
   - main関数はtask関数からheartbeatを受け取り、一定時間受け取らないとタイムアウトと判断する
   - timeoutはcontext.WithTimeoutによって判断する
*/

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// logging
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // errのチェックをしてからdeferで閉じる
	prefix := "ERROR: "
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), prefix, log.LstdFlags)

	// time management
	const mainTimeout = 5100 * time.Millisecond    // main関数の実行終了
	const beatInterval = 500 * time.Millisecond    // task関数のheartbeatの送信間隔
	const watchDogTimeout = 800 * time.Millisecond // task関数のheartbeatの受信タイムアウト

	ctx, cancel := context.WithTimeout(context.Background(), mainTimeout)
	defer cancel()

	// task関数の実行
	heartBeat, v := task(ctx, beatInterval)

	// watchDog timer

loop:
	for {
		select {
		case _, ok := <-heartBeat:
			// heartbeatチャネルの値がなくなったら終了
			if !ok {
				break loop
			}
			fmt.Println("⚡ pulse received ⚡")
		case r, ok := <-v:
			if !ok {
				break loop
			}
			time := strings.Split(r.String(), "m=")[1]
			fmt.Printf("★  value received ★: %s\n", time)
		case <-time.After(watchDogTimeout):
			errorLogger.Println("watchDog timer timed out. hartBeatの受信がありません。task関数が停止した可能性があります")
			break loop
		}
	}

}

/*
task関数は、下記を行う
1. haertbeatチャネルに、一定間隔で中身のない値を書き込む
2. outチャネルに、一定間隔で現在時刻を書き込む
3. ctxが終了したら、goroutineを終了する
*/
func task(
	ctx context.Context,
	beatInterval time.Duration,
) (<-chan struct{}, <-chan time.Time) {
	heartBeat := make(chan struct{}) // heartbeatの送信用チャネル. 値は不要なのでstruct{}型
	out := make(chan time.Time)      // 現在時刻の送信用チャネル

	go func() {
		defer close(heartBeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(beatInterval * 2)

		// haertBeatに値を書き込むための関数
		sendPulse := func() {
			select {
			case heartBeat <- struct{}{}:
			default:
			}
		}

		// 死活監視ためのheartbeatを送信し続ける
		// 現在時刻をoutに書き込む
		sendValue := func(t time.Time) {
			// returnされるまでループ
			for {
				select {
				case <-ctx.Done(): // そもそもctxが終了していたら抜ける
					return
				case <-pulse.C: // まずはpulseを優先. これ必要なのか？
					sendPulse()
				case out <- t: // pulseがなければ現在時刻をoutに書き込む
					return

				}
			}
		}

		// taskのメイン処理
		// var i int
		for {
			select {
			case <-ctx.Done(): // そもそもctxが終了していたら抜ける
				return
			case <-pulse.C:
				sendPulse()
				// 意図的に異常を起こす(あとで削除)
				// i++
				// if i > 3 {
				// 	time.Sleep(beatInterval * 3)
				// }
			case t := <-task.C: // 一定間隔で現在時刻を送信する
				sendValue(t)
			}
		}
	}()
	return heartBeat, out
}
