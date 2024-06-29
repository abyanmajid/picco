import axios from "axios";

import { client, endpoints } from "./api";

export async function getUserByEmail(email: string) {
  try {
    const response = await client.get(`${endpoints.user}/email/${email}`);

    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function createUser(username: string, email: string) {
  try {
    // Check if user already exists
    try {
      const existingUser = await getUserByEmail(email);

      if (existingUser) {
        // User already exists
        return {
          error: true,
          message: "User with this email already exists",
        };
      }
    } catch (error: any) {
      if (axios.isAxiosError(error)) {
        if (error.response && error.response.status === 404) {
          // User does not exist, we can proceed to create
        } else {
          // An unexpected error occurred
          throw error;
        }
      } else {
        // Non-Axios error occurred
        throw error;
      }
    }

    // Create new user
    const response = await client.post(endpoints.user, {
      username,
      email,
    });

    return {
      error: false,
      message: "User created successfully",
      data: response.data,
    };
  } catch (error: unknown) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        message: "Error creating user",
        details: error.response?.data,
      };
    } else {
      return {
        error: true,
        message: "An unexpected error occurred",
        details: error,
      };
    }
  }
}
