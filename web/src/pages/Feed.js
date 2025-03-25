import React from 'react';

export default function Feed() {
  return (
    <div className="container">
      <div className="grid">
        {Array.from({ length: 20 }, (_, index) => (
          <div key={index} className="grid-item">{index + 1}</div>
        ))}
      </div>
    </div>
  );
}