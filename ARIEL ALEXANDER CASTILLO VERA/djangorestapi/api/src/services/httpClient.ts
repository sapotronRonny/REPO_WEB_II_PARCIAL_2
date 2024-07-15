import axios, { AxiosInstance } from 'axios';
import config from './config';

class AxiosHttpClient {
  private axiosInstance: AxiosInstance;

  constructor() {
    this.axiosInstance = axios.create({
      baseURL: config.apiBaseUrl,  // URL base de tu API en Django
      timeout: 5000,  // Opcional: tiempo máximo de espera para las solicitudes en milisegundos
      headers: {
        'Content-Type': 'application/json',
        // Puedes añadir headers adicionales aquí según tus necesidades
      }
    });
  }

  async get<T>(url: string): Promise<T> {
    try {
      const response = await this.axiosInstance.get<T>(url);
      return response.data;
    } catch (error) {
      throw error;  // Maneja los errores según tus requisitos
    }
  }

  async post<T>(url: string, data: any): Promise<T> {
    try {
      const response = await this.axiosInstance.post<T>(url, data);
      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async put<T>(url: string, data: any): Promise<T> {
    try {
      const response = await this.axiosInstance.put<T>(url, data);
      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async delete<T>(url: string): Promise<T> {
    try {
      const response = await this.axiosInstance.delete<T>(url);
      return response.data;
    } catch (error) {
      throw error;
    }
  }
}

export default AxiosHttpClient;