import TwoValuesSwitch from "./components/TwoValuesSwitch";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import clsx from "clsx";

export default function BrowseProject() {
  const projectData = [[[{"Camera":{},"Filename":"VID_20240703_164016_00_002.insv","CameraPath":"Insta360","NormalizedTimestamp":"2024-07-03T14:40:16Z","LegacyTimestamp":"2024-07-03T14:40:16Z"},{"Camera":{},"Filename":"VID_20240703_164114_00_003.insv","CameraPath":"Insta360","NormalizedTimestamp":"2024-07-03T14:41:14Z","LegacyTimestamp":"2024-07-03T14:41:14Z"}],[{"Camera":{},"Filename":"DJI_0600.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T17:27:36Z","LegacyTimestamp":"2024-07-04T19:27:32Z"},{"Camera":{},"Filename":"DJI_0601.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T17:27:38Z","LegacyTimestamp":"2024-07-04T19:27:32Z"},{"Camera":{},"Filename":"DJI_0601.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T17:27:38Z","LegacyTimestamp":"2024-07-04T19:27:32Z"},{"Camera":{},"Filename":"DJI_0602.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T17:27:38Z","LegacyTimestamp":"2024-07-04T19:27:33Z"},{"Camera":{},"Filename":"DJI_0602.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T17:27:40Z","LegacyTimestamp":"2024-07-04T19:27:32Z"}],[{"Camera":{},"Filename":"DJI_0603.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:30:50Z","LegacyTimestamp":"2024-07-04T22:30:26Z"},{"Camera":{},"Filename":"DJI_0604.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:31:28Z","LegacyTimestamp":"2024-07-04T22:31:08Z"},{"Camera":{},"Filename":"DJI_0605.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:18Z","LegacyTimestamp":"2024-07-04T22:31:33Z"},{"Camera":{},"Filename":"DJI_0606.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:38Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0606.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:40Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0607.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:42Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0607.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:42Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0608.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:42Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0608.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:44Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0609.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:44Z","LegacyTimestamp":"2024-07-05T00:32:39Z"},{"Camera":{},"Filename":"DJI_0609.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:46Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0610.DNG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:46Z","LegacyTimestamp":"2024-07-05T00:32:38Z"},{"Camera":{},"Filename":"DJI_0610.JPG","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:32:46Z","LegacyTimestamp":"2024-07-05T00:32:39Z"},{"Camera":{},"Filename":"DJI_0611.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:33:20Z","LegacyTimestamp":"2024-07-04T22:33:01Z"},{"Camera":{},"Filename":"DJI_0612.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:34:10Z","LegacyTimestamp":"2024-07-04T22:33:25Z"},{"Camera":{},"Filename":"DJI_0613.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:34:54Z","LegacyTimestamp":"2024-07-04T22:34:33Z"},{"Camera":{},"Filename":"DJI_0614.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:35:14Z","LegacyTimestamp":"2024-07-04T22:34:57Z"},{"Camera":{},"Filename":"DJI_0615.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:35:54Z","LegacyTimestamp":"2024-07-04T22:35:41Z"},{"Camera":{},"Filename":"DJI_0616.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:37:04Z","LegacyTimestamp":"2024-07-04T22:36:20Z"},{"Camera":{},"Filename":"DJI_0617.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:37:52Z","LegacyTimestamp":"2024-07-04T22:37:15Z"},{"Camera":{},"Filename":"DJI_0618.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-03T22:38:40Z","LegacyTimestamp":"2024-07-04T22:38:12Z"}]],[[{"Camera":{},"Filename":"P1090342.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:05:29Z","LegacyTimestamp":"2024-07-04T17:05:29Z"},{"Camera":{},"Filename":"P1090343.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:05:59Z","LegacyTimestamp":"2024-07-04T17:05:59Z"},{"Camera":{},"Filename":"P1090344.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:25:15Z","LegacyTimestamp":"2024-07-04T17:25:15Z"},{"Camera":{},"Filename":"P1090345.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:27:11Z","LegacyTimestamp":"2024-07-04T17:27:11Z"},{"Camera":{},"Filename":"P1090346.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:38:28Z","LegacyTimestamp":"2024-07-04T17:38:28Z"},{"Camera":{},"Filename":"P1090347.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:38:46Z","LegacyTimestamp":"2024-07-04T17:38:46Z"},{"Camera":{},"Filename":"P1090348.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:52:09Z","LegacyTimestamp":"2024-07-04T17:52:09Z"},{"Camera":{},"Filename":"P1090349.MOV","CameraPath":"Lumix","NormalizedTimestamp":"2024-07-04T15:56:10Z","LegacyTimestamp":"2024-07-04T17:56:10Z"},{"Camera":{},"Filename":"DJI_0619.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:06:58Z","LegacyTimestamp":"2024-07-05T17:06:40Z"},{"Camera":{},"Filename":"DJI_0620.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:08:00Z","LegacyTimestamp":"2024-07-05T17:07:10Z"},{"Camera":{},"Filename":"DJI_0621.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:08:42Z","LegacyTimestamp":"2024-07-05T17:08:18Z"},{"Camera":{},"Filename":"DJI_0622.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:09:24Z","LegacyTimestamp":"2024-07-05T17:08:49Z"},{"Camera":{},"Filename":"DJI_0623.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:10:08Z","LegacyTimestamp":"2024-07-05T17:09:33Z"},{"Camera":{},"Filename":"DJI_0624.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:11:18Z","LegacyTimestamp":"2024-07-05T17:10:33Z"},{"Camera":{},"Filename":"DJI_0625.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:12:06Z","LegacyTimestamp":"2024-07-05T17:11:45Z"},{"Camera":{},"Filename":"DJI_0626.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:18:54Z","LegacyTimestamp":"2024-07-05T17:14:57Z"},{"Camera":{},"Filename":"DJI_0627.MP4","CameraPath":"Komarek","NormalizedTimestamp":"2024-07-04T16:20:48Z","LegacyTimestamp":"2024-07-05T17:19:02Z"},{"Camera":{},"Filename":"IMG_7331.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T16:33:53Z","LegacyTimestamp":"2024-07-04T16:33:53Z"},{"Camera":{},"Filename":"IMG_7332.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T16:33:54Z","LegacyTimestamp":"2024-07-04T16:33:54Z"},{"Camera":{},"Filename":"IMG_7331.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T16:33:55Z","LegacyTimestamp":"2024-07-04T16:33:55Z"},{"Camera":{},"Filename":"IMG_7332.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T16:33:55Z","LegacyTimestamp":"2024-07-04T16:33:55Z"}],[{"Camera":{},"Filename":"IMG_7333.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T17:06:24Z","LegacyTimestamp":"2024-07-04T17:06:24Z"},{"Camera":{},"Filename":"IMG_7334.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T17:14:49Z","LegacyTimestamp":"2024-07-04T17:14:49Z"}],[{"Camera":{},"Filename":"IMG_7335.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T17:49:46Z","LegacyTimestamp":"2024-07-04T17:49:46Z"},{"Camera":{},"Filename":"IMG_7336.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:07:02Z","LegacyTimestamp":"2024-07-04T18:07:02Z"},{"Camera":{},"Filename":"IMG_7337.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:07:30Z","LegacyTimestamp":"2024-07-04T18:07:30Z"},{"Camera":{},"Filename":"IMG_7338.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:01Z","LegacyTimestamp":"2024-07-04T18:08:01Z"},{"Camera":{},"Filename":"IMG_7339.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:02Z","LegacyTimestamp":"2024-07-04T18:08:02Z"},{"Camera":{},"Filename":"IMG_7338.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:03Z","LegacyTimestamp":"2024-07-04T18:08:03Z"},{"Camera":{},"Filename":"IMG_7339.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:03Z","LegacyTimestamp":"2024-07-04T18:08:03Z"},{"Camera":{},"Filename":"IMG_7340.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:19Z","LegacyTimestamp":"2024-07-04T18:08:19Z"},{"Camera":{},"Filename":"IMG_7341.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:20Z","LegacyTimestamp":"2024-07-04T18:08:20Z"},{"Camera":{},"Filename":"IMG_7340.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:21Z","LegacyTimestamp":"2024-07-04T18:08:21Z"},{"Camera":{},"Filename":"IMG_7341.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:21Z","LegacyTimestamp":"2024-07-04T18:08:21Z"},{"Camera":{},"Filename":"IMG_7342.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:08:34Z","LegacyTimestamp":"2024-07-04T18:08:34Z"},{"Camera":{},"Filename":"IMG_7343.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:33Z","LegacyTimestamp":"2024-07-04T18:09:33Z"},{"Camera":{},"Filename":"IMG_7343.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:34Z","LegacyTimestamp":"2024-07-04T18:09:34Z"},{"Camera":{},"Filename":"IMG_7344.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:35Z","LegacyTimestamp":"2024-07-04T18:09:35Z"},{"Camera":{},"Filename":"IMG_7344.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:36Z","LegacyTimestamp":"2024-07-04T18:09:36Z"},{"Camera":{},"Filename":"IMG_7345.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:38Z","LegacyTimestamp":"2024-07-04T18:09:38Z"},{"Camera":{},"Filename":"IMG_7345.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:38Z","LegacyTimestamp":"2024-07-04T18:09:38Z"},{"Camera":{},"Filename":"IMG_7346.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:46Z","LegacyTimestamp":"2024-07-04T18:09:46Z"},{"Camera":{},"Filename":"IMG_7347.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:46Z","LegacyTimestamp":"2024-07-04T18:09:46Z"},{"Camera":{},"Filename":"IMG_7346.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:47Z","LegacyTimestamp":"2024-07-04T18:09:47Z"},{"Camera":{},"Filename":"IMG_7347.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:09:48Z","LegacyTimestamp":"2024-07-04T18:09:48Z"},{"Camera":{},"Filename":"IMG_7348.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T18:27:52Z","LegacyTimestamp":"2024-07-04T18:27:52Z"}],[{"Camera":{},"Filename":"IMG_7349.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T19:01:18Z","LegacyTimestamp":"2024-07-04T19:01:18Z"},{"Camera":{},"Filename":"IMG_7349.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T19:01:18Z","LegacyTimestamp":"2024-07-04T19:01:18Z"}],[{"Camera":{},"Filename":"IMG_7350.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T19:51:59Z","LegacyTimestamp":"2024-07-04T19:51:59Z"},{"Camera":{},"Filename":"IMG_7350.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T19:51:59Z","LegacyTimestamp":"2024-07-04T19:51:59Z"}],[{"Camera":{},"Filename":"IMG_7351.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T20:51:28Z","LegacyTimestamp":"2024-07-04T20:51:28Z"},{"Camera":{},"Filename":"IMG_7351.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T20:51:30Z","LegacyTimestamp":"2024-07-04T20:51:30Z"}],[{"Camera":{},"Filename":"IMG_7357.MOV","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T23:54:08Z","LegacyTimestamp":"2024-07-04T23:54:08Z"},{"Camera":{},"Filename":"IMG_7358.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T23:55:27Z","LegacyTimestamp":"2024-07-04T23:55:27Z"},{"Camera":{},"Filename":"IMG_7359.HEIC","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T23:55:28Z","LegacyTimestamp":"2024-07-04T23:55:28Z"},{"Camera":{},"Filename":"IMG_7358.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T23:55:29Z","LegacyTimestamp":"2024-07-04T23:55:29Z"},{"Camera":{},"Filename":"IMG_7359.mov","CameraPath":"Tel Fil","NormalizedTimestamp":"2024-07-04T23:55:29Z","LegacyTimestamp":"2024-07-04T23:55:29Z"}]]];

  const [currentDay, setCurrentDay] = useState(0);
  const [currentBlock, setCurrentBlock] = useState(1);
  const [currentFileIndex, setCurrentFileIndex] = useState(0);
  const currentFile = projectData[currentDay][currentBlock][currentFileIndex];
  const [viewMode, setViewMode] = useState<'Tilemap' | 'Single preview'>('Tilemap');
  const [blockNames, setBlockNames] = useState<Record<string, string>>({});
  const [inputValue, setInputValue] = useState('');
  const navigate = useNavigate();
  
  // Update input value when switching blocks
  useEffect(() => {
    const blockKey = `${currentDay}-${currentBlock}`;
    setInputValue(blockNames[blockKey] || `Block ${currentBlock + 1}`);
  }, [currentDay, currentBlock, blockNames]);
  
  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const blockKey = `${currentDay}-${currentBlock}`;
    if (inputValue && inputValue.trim()) {
      setBlockNames(prev => ({
        ...prev,
        [blockKey]: inputValue.trim()
      }));
    }
  };

  function isVideo(filename: string) {
    return filename.endsWith('.MOV') || filename.endsWith('.mov') || filename.endsWith('.MP4');
  }
  
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
            <h2 className="text-xl font-semibold mb-0">Dzień {currentDay + 1}</h2>
            <h2 className="text-lg font-normal text-muted-foreground mb-4">Block {currentBlock + 1}</h2>
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
            {projectData[currentDay][currentBlock].map((file, index) => (
              <div key={file.CameraPath + file.Filename} className={clsx("transition border-[2px] w-[160px] h-[90px] bg-slate-500 flex-shrink-0 rounded-md overflow-hidden shadow-sm", file.Filename === currentFile.Filename && file.CameraPath === currentFile.CameraPath && file.NormalizedTimestamp === currentFile.NormalizedTimestamp ? 'border-red-500 shadow-md' : '')} onClick={() => setCurrentFileIndex(index)}>
              {isVideo(file.Filename) ? <div className="w-full h-full bg-slate-300">Video</div> : <img src={`/current-project/${file.CameraPath}/${file.Filename}`} alt={file.Filename} className="w-full h-full object-cover"  />}
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
            {projectData.map((_, dayIndex) => (
              <div 
                key={dayIndex}
                className={`w-[100px] h-[60px] flex-shrink-0 rounded-md overflow-hidden cursor-pointer border-2 ${
                  dayIndex === currentDay 
                    ? 'border-green-500 bg-green-50' 
                    : 'border-gray-200 hover:border-gray-300'
                }`}
                onClick={() => {
                  setCurrentDay(dayIndex);
                  setCurrentBlock(0); // Reset to first block when switching days
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
            {projectData[currentDay].map((_, blockIndex) => (
              <div 
                key={blockIndex}
                className={`w-[120px] h-[80px] flex-shrink-0 rounded-md overflow-hidden cursor-pointer border-2 ${
                  blockIndex === currentBlock 
                    ? 'border-blue-500 bg-blue-50' 
                    : 'border-gray-200 hover:border-gray-300'
                }`}
                onClick={() => {
                  setCurrentBlock(blockIndex);
                  setCurrentFileIndex(0); // Reset to first file when switching blocks
                }}
              >
                <div className="w-full h-full bg-gray-100 flex items-center justify-center">
                  <span className="text-sm font-medium text-gray-600 text-center px-1">
                    {blockNames[`${currentDay}-${blockIndex}`] || `Block ${blockIndex + 1}`}
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