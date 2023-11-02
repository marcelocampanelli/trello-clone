import axios from "axios";
import bearer from "$lib/user/bearerToken";
import isAuthorized from "$lib/user/isAuthorized";

axios.interceptors.response.use(isAuthorized);

const AxiosWithBearerApiV1 = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  timeout: 1000,
  headers: {
    "Content-Type": "application/json",
    Authorization: `Bearer ${bearer}`,
  },
});

export default AxiosWithBearerApiV1;
