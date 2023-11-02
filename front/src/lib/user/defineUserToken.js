const defineToken = async (value) => {
  localStorage.setItem("user_token", value);
};

export default defineToken;
