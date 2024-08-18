export const requiredRule = (msg, options) => {
  const defaults = {
    allowZero: false,
    allowEmptyString: false,
    isBoolean: false,
  };
  const { allowZero, allowEmptyString, isBoolean } = Object.assign(
    defaults,
    options
  );

  return (val) => {
    let errorMsg = typeof msg === "function" ? msg(val) : msg;
    errorMsg ??= "O campo acima é obrigatório.";

    if (typeof val === "string") {
      val = val.trim();
      return (allowEmptyString && val === "") || val.length > 0 || errorMsg;
    } else if (typeof val === "number") {
      return (allowZero && val === 0) || !!val || errorMsg;
    } else if (Array.isArray(val)) {
      return val.length > 0 || errorMsg;
    } else if (isBoolean === true) {
      return typeof val === "boolean" || errorMsg;
    }

    return !!val || errorMsg;
  };
};
