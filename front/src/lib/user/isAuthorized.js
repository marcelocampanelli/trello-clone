const isAuthorized = (response, config) => {
  if (response.status === 401) {
    window.location.href = "/login";
  }
};

export default isAuthorized;
