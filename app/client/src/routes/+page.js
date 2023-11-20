import { env } from "$env/dynamic/public";

const API_ENDPOINT = env.PUBLIC_API_ENDPOINT || "http://localhost:1323"

/** @type {import("./$types").PageLoad} */
export async function load({ fetch }) {
  const res = await fetch(API_ENDPOINT);
  const data = await res.json();

  return data;
}
