import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

const password = new random.RandomPet("password", {});
