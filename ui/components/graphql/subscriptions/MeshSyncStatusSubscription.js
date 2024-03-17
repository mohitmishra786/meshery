import { graphql, requestSubscription } from "react-relay";
import environment from "../../../lib/relayEnvironment";

const meshSyncStatusSubscription = graphql`
subscription MeshSyncStatusSubscription($k8scontextIDs: [String!]) {
  listenToMeshSyncEvents(k8scontextIDs: $k8scontextIDs) {
    contextID
    OperatorControllerStatus {
      name
      status
      version
      error {
        code
        description
      }
    }
  }
}
`;

export default function subscribeMeshSyncStatusEvents(dataCB, contextIds) {
  return requestSubscription(environment, {
    subscription : meshSyncStatusSubscription,
    variables : { k8scontextIDs : contextIds },
    onNext : dataCB,
    onError : (error) => console.log(`An error occured:`, error),
  });
}
