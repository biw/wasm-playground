// pollyfill
if (!WebAssembly.instantiateStreaming) {
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	};
}

window.GoWASM = {}

window._onGoWASMLoad = () => {
	console.log("called")
	if (window.onGoWASMLoad != null) {
		window.onGoWASMLoad(window.GOWASM)
	}
}

;(async () => {
	const go = new Go()
	const { instance } = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
	go.run(instance)
})()
