//@ts-nocheck

export function compileCode(src: string, value: any) {
  var initFunc = function (interpreter, globalObject) {
    const v = interpreter.nativeToPseudo(value);
    interpreter.setProperty(globalObject, "value", v);
  };
  const wrapper = `
  function run() {
    return ${src};
  }
  run();
  `;
  try {
    var myInterpreter = new window.JSInterpreter(wrapper, initFunc);
    let count = 0;
    while (true) {
      count++;
      if (count >= 500) {
        throw new Error(
          "Too many steps executed, possible infinite loop detected. A maximum of 500 steps is allowed."
        );
      }
      if (!myInterpreter.step()) {
        break;
      }
    }
    return myInterpreter.value;
  } catch (ex) {
    if (ex.message === "window.JSInterpreter is not a constructor") {
      alert("JS evalulator is not loaded, please try refreshing the page.");
      return;
    }
    throw ex;
  }
}
