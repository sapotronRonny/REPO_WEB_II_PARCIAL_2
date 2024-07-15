import json
from channels.generic.websocket import AsyncWebsocketConsumer
from .models import Pelicula

class YourConsumer(AsyncWebsocketConsumer):
    async def connect(self):
        await self.accept()
        await self.send_stats()

    async def disconnect(self, close_code):
        pass

    async def receive(self, text_data):
        await self.send_stats()

    async def send_stats(self):
        # Ejemplo de estadísticas: número de películas
        num_peliculas = Pelicula.objects.count()
        stats = {
            'num_peliculas': num_peliculas,
        }
        await self.send(text_data=json.dumps(stats))
