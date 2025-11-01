import axios from 'axios';

// The baseURL will be proxied by Nginx in Docker,
// but for local dev, it might point to localhost:8080
const api = axios.create({
  baseURL: '/api', // Use relative path for proxy
});

/**
 * Checks if vouchers exist for a flight.
 * @param {string} flightNumber
 * @param {string} date (YYYY-MM-DD)
 * @returns {Promise<object>} { exists: boolean }
 */
export const checkVoucherExists = (flightNumber, date) => {
  return api.post('/check', { flightNumber, date });
};

/**
 * Generates new vouchers.
 * @param {object} formData
 * @returns {Promise<object>} { success: boolean, seats: string[] }
 */
export const generateVouchers = (formData) => {
  return api.post('/generate', formData);
};