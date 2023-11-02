let tokenIsPresent = false;

if (localStorage.getItem("user_token") !== null) {
  tokenIsPresent = true;
}

export default tokenIsPresent;
