let tokenIsPresent = false;
let user = localStorage.getItem("user");


if (user !== undefined && user !== null) {
	tokenIsPresent = true;
}

export default tokenIsPresent;
