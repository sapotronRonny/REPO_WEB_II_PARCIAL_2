import AxiosHttpClient from '../httpClient';  // Ajusta la ruta según sea necesario
import config from '../config';

const httpClient = new AxiosHttpClient();

// Ejemplo de uso en una función asincrónica dentro de tu nuevo proyecto
async function obtenerDatos() {
  try {
    const url = `${config.apiBaseUrl}${config.endpoints.obtenerDatos}`;
    const data = await httpClient.get(url);
    console.log(data);
  } catch (error) {
    console.error('Error al obtener datos del servicio externo', error);
  }
}

obtenerDatos();