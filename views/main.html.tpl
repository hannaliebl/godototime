<h1>DOTOTIME</h1>
<table>
  <tbody>
    <tr>
      <th>ID</th>
      <th>Name</th>
      <th>State</th>
      <th>Game ID</th>
      <th>Game Name</th>
    </tr>
    {{range .Response.Players}}
      <tr>
        <td>{{.ID}}</td>
        <td>{{.Name}}</td>
        <td>{{.State}}</td>
        <td>{{.GameID}}</td>
        <td>{{.GameName}}</td>
      </tr>
    {{end}}
  </tbody>
</table>
