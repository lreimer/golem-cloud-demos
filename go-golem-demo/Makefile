build: compile
	wasm-tools component embed ./wit go_golem_demo.module.wasm --output go_golem_demo.embed.wasm
	wasm-tools component new go_golem_demo.embed.wasm -o go_golem_demo.wasm --adapt adapters/tier2/wasi_snapshot_preview1.wasm

bindings:
	wit-bindgen tiny-go --out-dir go_golem_demo ./wit

compile: bindings
	tinygo build -target=wasi -o go_golem_demo.module.wasm main.go

clean:
	rm -rf go_golem_demo
	rm *.wasm
