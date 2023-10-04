import React, { useState, useEffect } from "react";
import { API_URL } from "../config";
import { useParams, Link } from "react-router-dom";

export default function ViewUser() {
  let { id } = useParams();
  const [user, setUser] = useState({});
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchUser = async () => {
      setLoading(true);
      try {
        const response = await fetch(`${API_URL}/users/${id}`);
        const json = await response.json();
        setUser(json.data);
        setLoading(false);
      } catch (error) {
        console.log("error", error);
        setLoading(false);
      }
    };

    fetchUser();
  }, [id]);

  return (
    <div>
      {!loading ? (
        <div className="flex justify-center">
          <div className="lg:w-1/3 w-full">
            <div className="p-10">
              <div className="mb-10 flex items-center justify-between">
                <Link to="/">
                  <h1 className="font-bold">Go back</h1>
                </Link>
              </div>
              <div className="bg-slate-200 rounded-lg">
                <div className="flex border-b py-4 pl-2">
                  <div className="mr-4 text-slate-400">Name</div>
                  <div className="text-slate-800">{user.name}</div>
                </div>
                <div className="flex border-b py-4 pl-2">
                  <div className="mr-4 text-slate-400">Email</div>
                  <div className="text-slate-800">{user.email}</div>
                </div>
                <div className="flex border-b py-4 pl-2">
                  <div className="mr-4 text-slate-400">Date</div>
                  <div className="text-slate-800">{user.date}</div>
                </div>
                <div className="flex border-b py-4 pl-2">
                  <div className="mr-4 text-slate-400">City</div>
                  <div className="text-slate-800">{user.city}</div>
                </div>
                <div className="flex border-b py-4 pl-2">
                  <div className="mr-4 text-slate-400">Country</div>
                  <div className="text-slate-800">{user.country}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      ) : (
        ""
      )}
    </div>
  );
}
