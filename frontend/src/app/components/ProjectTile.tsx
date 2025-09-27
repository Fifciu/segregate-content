import {
  Card,
  // CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

interface ProjectTileProps {
  year: string
  title: string
  description?: string
  description2?: string
}

export default function ProjectTile({ year, title, description = 'Trending up this month', description2 = 'Visitors for the last 6 months' }: ProjectTileProps) {
  return (
    <Card className="@container/card">
        <CardHeader>
          <CardDescription>{year}</CardDescription>
          <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
            {title}
          </CardTitle>
          {/* <CardAction>
            <Badge variant="outline">
              <IconTrendingUp />
              +12.5%
            </Badge>
          </CardAction> */}
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium">
            {description}
          </div>
          <div className="text-muted-foreground">
            {description2}
          </div>
        </CardFooter>
      </Card>
  )
}