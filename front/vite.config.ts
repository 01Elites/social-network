import devtools from 'solid-devtools/vite';
import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
import tsconfigPaths from 'vite-tsconfig-paths';

// for github pages
const isRelativeBuild = process.env.RELATIVE_BUILD === 'true';

export default defineConfig({
  base: isRelativeBuild ? './' : '/',
  plugins: [
    tsconfigPaths(),
    /* 
    Uncomment the following line to enable solid-devtools.
    For more info see https://github.com/thetarnav/solid-devtools/tree/main/packages/extension#readme
    */
    devtools({
      locator: {
        jsxLocation: true,
        componentLocation: true,
      },
    }),
    solidPlugin(),
  ],
  server: {
    port: 3000,
    host: '0.0.0.0',
  },
  build: {
    target: 'esnext',
  },
});
