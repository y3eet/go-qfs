import type { FileType } from "../types";
import { getFiles } from "../utils/api";

export const fileStore = $state({
  files: [] as FileType[],
  fetchLoading: false,
});

export async function fetchFiles(dir: string[]) {
  fileStore.fetchLoading = true;
  fileStore.files = await getFiles(dir.join("/"));
  fileStore.fetchLoading = false;
}
