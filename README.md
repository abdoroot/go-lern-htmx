Golang-htmx-tailwind Starter with hot reloading
#export go/bin folder : export PATH=$PATH:/home/abdoroot/go/bin
#hotreload using air air -c .air.toml #air only
watch both templ and air:
 generate --watch --proxy="http://localhost:3001" --cmd="air air -c .air.tom"