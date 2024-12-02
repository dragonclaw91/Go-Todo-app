import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { create } from 'zustand'
import View from "./client.jsx"


export const useInputStore = create((set) => ({
    inputValue: '',
    updateInputValue: (newValue) => set({ inputValue: newValue }),
    log: console.log("VALUE")
  }));
  


export const  useStore = create((set) => ({
    data: [],
    isLoading: "TEST",
    deleteData: async (id) => {
        console.log("DELETEING CONTROLLER",id)
        try {
          set({ loading: true, error: null });
    
          const response = await   axios.delete(`http://localhost:5000/${id}`)
  
         console.log("DELETE RESPONSE",response)
          if (!response) {
            throw new Error("Failed to DELETE data");
          }
    
          // Optional: return response or status to caller
        //   return await response.json();
        } catch (error) {
          set({ error: error.message });
          throw error; // Rethrow error for caller to handle
        } finally {
          set({ loading: false });
        }
      },
    updateData: async (id) => {
        try {
          set({ loading: true, error: null });
    
          const response = await   axios.put('http://localhost:5000', {"id":`${id}`})
  
         
          if (!response) {
            throw new Error("Failed to POST data");
          }
    
          // Optional: return response or status to caller
        //   return await response.json();
        } catch (error) {
          set({ error: error.message });
          throw error; // Rethrow error for caller to handle
        } finally {
          set({ loading: false });
        }
      },
    postData: async (id) => {
        try {
          set({ loading: true, error: null });
    
          const response = await   axios.post('http://localhost:5000', {"task":`${id}`})
  
         
          if (!response) {
            throw new Error("Failed to POST data");
          }
    
          // Optional: return response or status to caller
        //   return await response.json();
        } catch (error) {
          set({ error: error.message });
          throw error; // Rethrow error for caller to handle
        } finally {
          set({ loading: false });
        }
      },
    fetchData: async() => {
        set({ isLoading: true });
        try{
            console.log("TRYING")
      const res = await axios.get('http://localhost:5000')
       
    console.log("RESPONSE",res)
        const data =  res
        
        set({ data, isLoading: false });
    }
        catch (error) {
            set({ error, isLoading: false });
          }
    }
  }))







  

