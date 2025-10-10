import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function TwoValuesSwitch({ 
  optionA,
  optionB,
  value,
  onChange
}: { optionA: string, optionB: string, value: string,onChange: (value: string) => void }) {
  const [val, setVal] = useState(value || optionA);
  function onClick(value: string) {
    setVal(value);
    onChange(value);
  }

  return (
    <div className="flex">
      <Button variant={val === optionA ? 'default' : 'outline'} className="px-4 py-2 rounded-[0px] rounded-l-md" onClick={() => onClick(optionA)}>{optionA}</Button>
      <Button variant={val === optionB ? 'default' : 'outline'} className="px-4 py-2 rounded-[0px] rounded-r-md" onClick={() => onClick(optionB)}>{optionB}</Button>
    </div>
  )
}