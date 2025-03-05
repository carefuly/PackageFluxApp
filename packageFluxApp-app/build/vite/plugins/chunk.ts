import {chunkSplitPlugin} from 'vite-plugin-chunk-split'

export default function createChunkSplit() {
  return chunkSplitPlugin({
    strategy: 'default',
  })
}
