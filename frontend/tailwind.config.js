/** @type {import('tailwindcss').Config} */
import daisyui from 'daisyui'
import preline from 'preline/plugin';
import flowbite from 'flowbite/plugin';

import tforms from "@tailwindcss/forms";
import ttypo from "@tailwindcss/typography";

export default {
  content: [
    'node_modules/preline/dist/*.js',
    'templates/**/*.templ',
    //"./node_modules/flowbite/**/*.js"
  ],
  darkMode: 'class',
  /*daisyui: {
    themes: ["light", "dark", "nihgt"],
  },*/
  theme: {
    colors: {
      'blue': '#1fb6ff',
      'purple': '#7e5bef',
      'pink': '#ff49db',
      'orange': '#ff7849',
      'green': '#13ce66',
      'yellow': '#ffc82c',
      'gray-dark': '#273444',
      'gray': '#8492a6',
      'gray-light': '#d3dce6',
    },
    fontFamily: {
      sans: ['Graphik', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
    },
    extend: {
      spacing: {
        '8xl': '96rem',
        '9xl': '128rem',
      },
      borderRadius: {
        '4xl': '2rem',
      }
    }
  },
  plugins: [
    tforms,
    ttypo,
    preline
    //daisyui
    //require('flowbite/plugin')
    //require('@tailwindcss/forms'),
    //require('preline/plugin'),
  ],
}

