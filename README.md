To run this code, you have to run 2 programs at the same time :

Execute (from the root of the git repo) :
	go run server/src/app/main.go
You can check on your browser
http://localhost:8083/ping 
It should display "Server is up!"

Then, execute :
	cd client
	npm run dev

You can now find the web-app on http://localhost:3000/ on your browser !