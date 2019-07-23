// AUTH_TOKEN localStorage name for the authentication token
export const AUTH_TOKEN = 'auth-token';

// API_ENDPOINT is the API endpoint for production and developement
export const API_ENDPOINT = process.env.NODE_ENV === 'production' ? '' : 'http://localhost:8081';
