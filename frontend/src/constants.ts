//export const API_URL = "https://backend-production-cd54.up.railway.app"
export const API_URL =
  process.env.NODE_ENV === "development"
    ? "http://localhost:2680/api"
    : "https://backend-production-cd54.up.railway.app/api";
