interface Config {
  API_URL: string;
  WS_URL: string;
}

const config: Config = {
  API_URL: 'http://localhost:8080/api',
  WS_URL: 'ws://localhost:8080/api/ws', // WebSocket URL
};

export default config;
