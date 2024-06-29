import axios from "axios";

const brokerEndpoint = process.env.BROKER_ENDPOINT;

export const endpoints = {
  user: `${brokerEndpoint}/user`,
  notification: `${brokerEndpoint}/notification`,
  mail: `${brokerEndpoint}/mail`,
  judge: `${brokerEndpoint}/judge`,
  compiler: `${brokerEndpoint}/compiler`,
  course: `${brokerEndpoint}/course`,
};

export const client = axios.create({
  baseURL: brokerEndpoint,
  headers: {
    "Content-Type": "application/json",
  },
});
