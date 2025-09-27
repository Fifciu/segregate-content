import {
  Card,
  // CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Link } from "react-router"

export default function NewProjectTile() {
  return (
    <Card className="@container/card bg-blue-200">
      <CardHeader>
        <CardDescription>Kliknij aby dodać</CardDescription>
        <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
          Nowy projekt
        </CardTitle>
        {/* <CardAction>
            <Badge variant="outline">
              <IconTrendingUp />
              +12.5%
            </Badge>
          </CardAction> */}
      </CardHeader>
      <CardFooter className="flex w-full justify-between">
        <div className="flex-col items-start gap-1.5 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium">
            Usprawnij swoją pracę
          </div>
          <div className="text-muted-foreground">
            Szybko i prosto
          </div>
        </div>
        <div>
          <Button asChild>
            <Link to="/new-project">Dodaj projekt</Link>
          </Button>
        </div>
      </CardFooter>
    </Card>
  )
}