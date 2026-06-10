import type { FileType } from "../types";
const BASE_URL = "http://localhost:8080";

export async function getFiles() {
  return await fetch(`${BASE_URL}/api/files`).then(
    (res) => res.json() as Promise<FileType[]>,
  );
}
