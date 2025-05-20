import { API_URL } from "./constants";

export async function createDocument(documentTitle: string) {
  const res = await fetch(API_URL + "/api/document/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title: documentTitle }),
  });
  console.log(await res.json());
}