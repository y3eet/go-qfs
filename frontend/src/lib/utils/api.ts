import type { FileType } from "../types";
const BASE_URL = "";

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

export async function uploadFile(file: File, path: string) {
  return fetch(`${BASE_URL}/api/file/upload?path=${encodeURIComponent(path)}`, {
    method: "POST",
    body: file,
    headers: {
      "Content-Type": file.type,
      "X-Filename": file.name,
    },
  });
}
