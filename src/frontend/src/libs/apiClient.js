import axios from "axios";

const ipAddr = import.meta.env.VITE_IPADDR;

const apiClient = axios.create({
  baseURL: `http://${ipAddr}:8080`,
});

export default apiClient;
