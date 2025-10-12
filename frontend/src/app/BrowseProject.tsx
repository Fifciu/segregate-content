import TwoValuesSwitch from "./components/TwoValuesSwitch";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import clsx from "clsx";
import { useBlocksData } from "./hooks/useBlocksData.tsx";

export default function BrowseProject() {
  const { days, setBlock } = useBlocksData();
  if(!days.length) {
    return <div>No days</div>;
  }
  const [currentDayIndex, setCurrentDayIndex] = useState(0);
  const [currentBlockIndex, setCurrentBlockIndex] = useState(Object.keys(days[currentDayIndex]?.blocks)?.[0] || '');
  const [currentFileIndex, setCurrentFileIndex] = useState(0);
  const currentDay = days[currentDayIndex];
  const currentBlock = currentDay.blocks[currentBlockIndex];
  const currentFile = currentBlock.files[currentFileIndex];
  const [viewMode, setViewMode] = useState<'Tilemap' | 'Single preview'>('Tilemap');
  const [inputValue, setInputValue] = useState('');
  const navigate = useNavigate();
  
  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (inputValue && inputValue.trim()) {
      setBlock(currentBlockIndex, currentDayIndex, inputValue.trim());
    }
  };

  function isVideo(filename: string) {
    return filename.endsWith('.MOV') || filename.endsWith('.mov') || filename.endsWith('.MP4');
  }

  useEffect(() => {
    setInputValue(currentBlock.name);
  }, [currentDayIndex, currentBlockIndex, currentFileIndex])
  
  return (
    <div className="container mx-auto py-8 px-4 max-w-full overflow-hidden">
      {/* <div className="mb-6">
        <h1 className="text-3xl font-bold">Szczegóły projektu</h1>
        <p className="text-muted-foreground mt-2">
          Dostosuj efekty wstępnego podziału grup
        </p>
      </div> */}
      
      <div className="flex gap-6 w-full max-w-full overflow-hidden">
        {/* Form Section - 60% width */}
        <div className="w-[60%] min-w-0">
          <div className="grid grid-cols-3 justify-between">
            <div>
            <h2 className="text-xl font-semibold mb-0">Dzień {currentDayIndex + 1}</h2>
            <h2 className="text-lg font-normal text-muted-foreground mb-4">Block {currentBlockIndex + 1}</h2>
            </div>
            <div className="justify-self-center w-[220px]">
              <TwoValuesSwitch optionA="Tilemap" optionB="Single preview" value={viewMode} onChange={(val) => setViewMode(val as any)}/>
            </div>
          </div>
        { currentFile && 
          <div className="space-y-2 rounded-md overflow-hidden object-contain h-[518px] bg-slate-400">
            <div className="h-[100%] w-fit mx-auto shadow-xl">
              {isVideo(currentFile.Filename) ? <video src={`/current-project/${currentFile.CameraPath}/${currentFile.Filename}`} controls className="h-[100%]"/> : <img src={`/current-project/${currentFile.CameraPath}/${currentFile.Filename}`} alt={currentFile.Filename} className="h-[100%]"/>}
            </div>
            {/* Add your content here */}
          </div>
        }
          <div className="flex overflow-x-auto gap-2 mt-2 pb-2 h-[112px]">
            {currentBlock.files.map((file, index) => (
              <div key={file.CameraPath + file.Filename} className={clsx("transition border-[2px] w-[160px] h-[90px] bg-slate-500 flex-shrink-0 rounded-md overflow-hidden shadow-sm", file.Filename === currentFile.Filename && file.CameraPath === currentFile.CameraPath && file.NormalizedTimestamp === currentFile.NormalizedTimestamp ? 'border-red-500 shadow-md' : '')} onClick={() => setCurrentFileIndex(index)}>
              {isVideo(file.Filename) ? <img src={`/current-project/thumbnails/${file.CameraPath}/${file.Filename}`} alt={file.Filename} className="w-full h-full object-cover"  /> : <img src={`/current-project/${file.CameraPath}/${file.Filename}`} alt={file.Filename} className="w-full h-full object-cover"  />}
              </div>
            ))}
          </div>
        </div>
      
      {/* Vertical divider */}
      <div className="w-px bg-gray-300 flex-shrink-0"></div>
      
      {/* Second section - remaining width */}
      <div className="flex-1 p-4 flex flex-col min-w-0 max-w-full">
        <h2 className="text-xl font-semibold mb-4">Details</h2>
        
        {/* Block name form */}
        <div className="mb-6">
          <form onSubmit={handleFormSubmit}>
            <div className="flex gap-3 items-end">
              <div className="flex-1">
                <label htmlFor="blockName" className="block text-sm font-medium text-gray-700 mb-2">
                  Block name
                </label>
                <input
                  type="text"
                  id="blockName"
                  name="blockName"
                  className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="Enter block name"
                  value={inputValue}
                  onChange={(e) => setInputValue(e.target.value)}
                />
              </div>
              <button
                type="submit"
                className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 whitespace-nowrap"
              >
                Save
              </button>
            </div>
          </form>
        </div>

        {/* Day carousel */}
        <div className="mt-auto mb-6 min-w-0">
          <h3 className="text-lg font-semibold mb-3">Days</h3>
          <div className="flex overflow-x-auto gap-2 pb-2">
            {days.map((_, dayIndex) => (
              <div 
                key={dayIndex}
                className={`w-[100px] h-[60px] flex-shrink-0 rounded-md overflow-hidden cursor-pointer border-2 ${
                  dayIndex === currentDayIndex 
                    ? 'border-green-500 bg-green-50' 
                    : 'border-gray-200 hover:border-gray-300'
                }`}
                onClick={() => {
                  setCurrentDayIndex(dayIndex);
                  setCurrentBlockIndex(Object.keys(days[dayIndex].blocks)[0]); // Reset to first block when switching days
                  setCurrentFileIndex(0); // Reset to first file when switching days
                }}
              >
                <div className="w-full h-full bg-gray-100 flex items-center justify-center">
                  <span className="text-sm font-medium text-gray-600">
                    Day {dayIndex + 1}
                  </span>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Block carousel */}
        <div className="min-w-0">
          <h3 className="text-lg font-semibold mb-3">Blocks</h3>
          <div className="flex overflow-x-auto gap-2 pb-2 h-[102px]">
            {Object.entries(currentDay.blocks).map(([blockUuid, block]) => (
              <div 
                key={blockUuid}
                className={`w-[120px] h-[80px] flex-shrink-0 rounded-md overflow-hidden cursor-pointer border-2 ${
                  blockUuid === currentBlockIndex 
                    ? 'border-blue-500 bg-blue-50' 
                    : 'border-gray-200 hover:border-gray-300'
                }`}
                onClick={() => {
                  setCurrentBlockIndex(blockUuid);
                  setCurrentFileIndex(0); // Reset to first file when switching blocks
                }}
              >
                <div className="w-full h-full bg-gray-100 flex items-center justify-center">
                  <span className="text-sm font-medium text-gray-600 text-center px-1">
                    {block.name}
                  </span>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Go to summary button */}
        <div className="mt-6 flex justify-end">
          <button
            onClick={() => navigate('/project-summary')}
            className="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          >
            Go to summary
          </button>
        </div>
        </div>
      </div>
    </div>
  );
}