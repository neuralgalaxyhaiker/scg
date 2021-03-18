build: clean
	@mkdir -p dist
	go mod download
	go build -o dist/lancher

encode: build
	./dist/lancher encode ../nodejs/dist/server.js ./dist/server.out.js

clean:
	@rm -rf dist

dist: encode
	@cp ../nodejs/package.json dist/
	@cp -r ../nodejs/node_modules dist/

dev: nodedep
	./dist/lancher ./dist/server.out.js