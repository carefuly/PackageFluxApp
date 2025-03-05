import { compression } from 'vite-plugin-compression2';
export default function createCompression(env, isBuild) {
    const plugin = [];
    if (isBuild) {
        const { VITE_BUILD_COMPRESS } = env;
        const compressList = VITE_BUILD_COMPRESS.split(',');
        if (compressList.includes('gzip')) {
            plugin.push(compression());
        }
        if (compressList.includes('brotli')) {
            plugin.push(compression({
                exclude: [/\.(br)$/, /\.(gz)$/],
                algorithm: 'brotliCompress',
            }));
        }
    }
    return plugin;
}
//# sourceMappingURL=compression.js.map