import type { FileType } from "../types";
const BASE_URL = "http://localhost:8080";

export async function getFiles(path: string) {
  return await fetch(
    `${BASE_URL}/api/files?path=${encodeURIComponent(path)}`,
  ).then((res) => res.json() as Promise<FileType[]>);
}

export async function downloadFile(file: string) {
  const res = await fetch(`${BASE_URL}/api/file/download/${file}`);
  if (!res.ok) throw new Error("Download failed");

  const blob = await res.blob();
  const url = URL.createObjectURL(blob);

  const a = document.createElement("a");
  a.href = url;
  a.download = file;
  a.click();

  URL.revokeObjectURL(url);
}
