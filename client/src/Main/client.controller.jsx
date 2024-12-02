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
    fetchData: async() => {
        set({ isLoading: true });
        try{
            console.log("TRYING")
      const res = await axios.get('http://localhost:5000')
        .then(response => {
          console.log("FETCHING")
         return response
        })
        .catch(error => {
          console.log(error);
        });
    console.log("RESPONSE",res)
        const data =  res
        
        set({ data, isLoading: false });
    }
        catch (error) {
            set({ error, isLoading: false });
          }
    }
  }))

 export async function post(input){ axios.post('http://localhost:5000', input)
  .then(function (response) {
    console.log("POSTED",response, View.fetchData);
   
    // useStore()
  })
  .catch(function (error) {
    console.log("POSTED",error);
  });
 }

 export function put(input){ axios.put('http://localhost:5000', input)
 .then(function (response) {
   console.log("PUTED");
 })
 .catch(function (error) {
   console.log("PUTED",error);
 });
}

  

