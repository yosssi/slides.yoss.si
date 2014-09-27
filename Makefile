bin:
	go-bindata -o ./assets.go public/... slides/...
run:
	go run main.go assets.go slides-yoss-si
runp:
	MARTINI_ENV=production go run main.go assets.go slides-yoss-si
stop:
	ps aux | grep slides-yoss-si | grep exe/main | grep -v grep | awk '{print $$2}' | xargs kill -9
rerun:
	make stop
	make bin
	make run
themes:
	grunt themes
