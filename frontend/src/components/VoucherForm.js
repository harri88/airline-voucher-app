import React, { useState } from 'react';
import { checkVoucherExists, generateVouchers } from '../services/api';

const VoucherForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    id: '',
    flightNumber: '',
    date: '',
    aircraft: 'ATR', // Default value
  });

  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [generatedSeats, setGeneratedSeats] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);
    setGeneratedSeats(null);

    try {
      // 1. Check if vouchers already exist
      const checkResponse = await checkVoucherExists(
        formData.flightNumber,
        formData.date
      );

      if (checkResponse.data.exists) {
        setError('Vouchers have already been generated for this flight and date.');
        setIsLoading(false);
        return;
      }

      // 2. If not, generate new vouchers
      const genResponse = await generateVouchers(formData);

      if (genResponse.data.success) {
        setGeneratedSeats(genResponse.data.seats);
      }
    } catch (err) {
      if (err.response && err.response.data && err.response.data.error) {
        setError(err.response.data.error);
      } else {
        setError('An unexpected error occurred. Please try again.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="form-container">
      <h2>Airline Voucher Seat Assignment</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Crew Name</label>
          <input
            type="text"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label>Crew ID</label>
          <input
            type="text"
            name="id"
            value={formData.id}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label>Flight Number</label>
          <input
            type="text"
            name="flightNumber"
            value={formData.flightNumber}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label>Flight Date</label>
          <input
            type="date"
            name="date"
            value={formData.date}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label>Aircraft Type</label>
          <select
            name="aircraft"
            value={formData.aircraft}
            onChange={handleChange}
            required
          >
            <option value="ATR">ATR</option>
            <option value="Airbus 320">Airbus 320</option>
            <option value="Boeing 737 Max">Boeing 737 Max</option>
          </select>
        </div>
        <button type="submit" disabled={isLoading}>
          {isLoading ? 'Generating...' : 'Generate Vouchers'}
        </button>
      </form>

      {error && <div className="message error">{error}</div>}

      {generatedSeats && (
        <div className="message success">
          <h3>Vouchers Generated Successfully!</h3>
          <p>Assigned Seats: <strong>{generatedSeats.join(', ')}</strong></p>
        </div>
      )}
    </div>
  );
};

export default VoucherForm;