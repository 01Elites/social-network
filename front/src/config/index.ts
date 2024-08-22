interface Config {
  API_URL: string;
  WS_URL: string;
}

const config: Config = {
  API_URL: `http://${window.location.host}/api`,
  WS_URL: `ws://${window.location.host}/api/ws`, // WebSocket URL
};

export default config;
