

// our program needs to be in a callback in order to make sure that the
// wasm files have been loaded in async
// for convencice, the global window.GoWASM is passed as a param into the func
const main = () => {
	const { goPrint } = window.GoWASM

	console.log('bang')
	goPrint('bang from golang')
}

function setRender () {}



window.onGoWASMLoad = main
