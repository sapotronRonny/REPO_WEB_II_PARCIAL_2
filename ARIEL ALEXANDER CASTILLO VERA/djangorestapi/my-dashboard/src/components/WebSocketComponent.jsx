import React, { useEffect, useState } from 'react';
import ReconnectingWebSocket from 'reconnecting-websocket';

const WebSocketComponent = () => {
    const [stats, setStats] = useState({ num_peliculas: 0 });

    useEffect(() => {
        const rws = new ReconnectingWebSocket('ws://localhost:8000/ws/somepath/');

        rws.onopen = () => {
            console.log('WebSocket connection established');
        };

        rws.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setStats(data);
        };

        rws.onclose = () => {
            console.log('WebSocket connection closed');
        };

        rws.onerror = (error) => {
            console.log('WebSocket error: ' + error);
        };

        return () => {
            rws.close();
        };
    }, []);

    return (
        <div>
            <h1>WebSocket Dashboard</h1>
            <p>Number of Movies: {stats.num_peliculas}</p>
        </div>
    );
};

export default WebSocketComponent;
