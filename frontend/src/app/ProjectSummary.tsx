import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Link } from "react-router";
import clsx from "clsx";
import { useState } from "react";
import { useBlocksData } from "./hooks/useBlocksData.tsx";

export default function ProjectSummary() {
  const { days } = useBlocksData();

  const fromDirectories = days.reduce((total, day) => {
    for (let block of Object.values(day.blocks)) {
      for (let file of block.files) {
        if (!total[file.CameraPath]) {
          total[file.CameraPath] = [];
        }
        total[file.CameraPath].push(file.Filename);
      }
    }
    return total
  }, {} as Record<string, string[]>);

  return (
    <div className="container mx-auto py-8 px-4">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">Podsumowanie projektu</h1>
        <p className="text-muted-foreground mt-2">
          Przegląd i analiza segregacji treści
        </p>
      </div>
      <div className="grid w-full grid-cols-[2fr_auto_5fr] gap-5">
        <div>
          <Card>
            <CardHeader>
              <CardTitle>Obecna struktura plików</CardTitle>
              <CardDescription>
                Katalogi w projekcie
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                {Object.keys(fromDirectories).map((dir) => (
                  <Directory name={dir as string} files={fromDirectories[dir as string]} key={dir as string} />
                ))}
              </div>
            </CardContent>
          </Card>
        </div>
        <div className="w-px bg-gray-300 flex-shrink-0"></div>
        <div className="max-h-[70vh] overflow-y-auto mr-1">
          <Card>
            <CardHeader>
              <CardTitle>Zbudowana struktura plików</CardTitle>
              <CardDescription>
                Katalogi w projekcie
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                {days.map((day, index) => {
                  return (
                    <>
                      <Directory name={`Dzień ${index + 1}`} key={index} />
                      <div className="pl-2">
                        {Object.entries(day.blocks).map(([uuid, block]) => <Directory name={block.name} key={uuid} files={block.files.map(b => b.Filename)} />)}
                      </div>
                    </>
                  )
                })}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
      {/* Action Buttons */}
      <div className="mt-8 flex justify-between gap-4">
        <Button asChild variant="secondary">
          <Link to="/browse-project/1">Cofnij do przeglądu</Link>
        </Button>
        <Button asChild variant="outline">
          <Link to="/">Home</Link>
        </Button>
        <Button asChild>
          <Link to="/copying-files">Ustaw kopiowanie</Link>
        </Button>
      </div>
    </div>
  );
}

function Directory({ name, files }: { name: string, files?: string[] }) {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div>
      <div className="flex justify-between">
        <h3 className="text-lg font-semibold mb-3">{name}</h3>
        {files && <Button variant="outline" onClick={() => setIsOpen(!isOpen)}>{isOpen ? 'Hide' : 'View'}</Button>}
      </div>
      {files && <div className={clsx("my-2 grid transition-[grid-template-rows] duration-300 max-h-[200px]", isOpen ? "grid-rows-[1fr] overflow-y-auto" : "grid-rows-[0fr] overflow-hidden")}>
        <div className={clsx("min-h-0 pl-2")}>
          {files.map((file) => (
            <div key={file}>{file}</div>
          ))}
        </div>
      </div>}
    </div>
  )
}