import {client, endpoints} from "./api"
import axios from "axios";
import { AxiosError } from "axios";

export async function getCourse(title: string) {
  try {
    const response = await client.get(`${endpoints.course}/${title}`);

    return response.data.data;
  } catch (error: AxiosError | unknown) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error response:", error.response);
    } else {
      console.error("Unknown error occurred:", error);
    }
    throw error;
  }
}
