import axios from "axios";

const axiosApiV1 = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  timeout: 1000,
  headers: {
    "Content-Type": "application/json",
  },
});

export default axiosApiV1;
