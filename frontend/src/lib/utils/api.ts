import type { FileType } from "../types";
const BASE_URL = "http://localhost:8080";

export async function getFiles(path: string) {
  return await fetch(
    `${BASE_URL}/api/files?path=${encodeURIComponent(path)}`,
  ).then((res) => res.json() as Promise<FileType[]>);
}

export async function downloadFile(filePath: string) {
  const cleanPath = filePath.replace(/^\//, "");
  const encodedPath = cleanPath.split("/").map(encodeURIComponent).join("/");
  window.location.href = `${BASE_URL}/api/file/download/${encodedPath}`;
}
