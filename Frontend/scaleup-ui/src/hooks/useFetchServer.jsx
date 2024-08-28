import { useState, useEffect } from "react";

const useFetchServers = () => {
  const [servers, setServers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchServerData = async () => {
      try {
        const resp = await fetch("/api/servers");
        if (!resp.ok) {
          throw new Error("Failed to fetch data");
        }
        const data = await resp.json();
        setServers(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchServerData();
  }, []);

  return { servers, loading, error };
};

export default useFetchServers;
