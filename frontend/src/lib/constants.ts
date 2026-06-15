import {
  FileText,
  Image,
  Headphones,
  Play,
  FileArchive,
  File,
  Folder,
} from "@lucide/svelte";
import type { Component } from "svelte";

export const MIMES: Record<string, { type: string; icon: Component }> = {
  //Document & Text Formats
  ".txt": {
    type: "text/plain",
    icon: FileText,
  },
  ".md": {
    type: "text/markdown",
    icon: FileText,
  },
  ".pdf": {
    type: "application/pdf",
    icon: FileText,
  },

  // Image Formats
  ".jpg": {
    type: "image/jpeg",
    icon: Image,
  },
  ".jpeg": {
    type: "image/jpeg",
    icon: Image,
  },
  ".png": {
    type: "image/png",
    icon: Image,
  },
  ".gif": {
    type: "image/gif",
    icon: Image,
  },
  ".svg": {
    type: "image/svg+xml",
    icon: Image,
  },
  ".webp": {
    type: "image/webp",
    icon: Image,
  },

  //  Media & Audio Formats
  ".mp3": {
    type: "audio/mpeg",
    icon: Headphones,
  },
  ".mp4": {
    type: "video/mp4",
    icon: Play,
  },

  //  Archive Formats
  ".zip": {
    type: "application/zip",
    icon: FileArchive,
  },
  ".tar": {
    type: "application/x-tar",
    icon: FileArchive,
  },
  ".gz": {
    type: "application/gzip",
    icon: FileArchive,
  },

  // Unknown Formats
  default: {
    type: "application/octet-stream",
    icon: File,
  },

  dir: {
    type: "directory",
    icon: Folder,
  },
};
