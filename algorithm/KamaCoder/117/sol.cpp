#include <bits/stdc++.h>
using namespace std;

const int N = 1e5 + 10;

int n, m;

int main() {
    cin >> n >> m;
    cout << "n = " << n << "m = " << m << endl;
    vector<vector<int>> graph(n + 1, vector<int> (n + 1));
    vector<int> in(n + 1);
    vector<int> out(n + 1);
    int u, v;
    for (int i = 0 ; i < m ; i ++) {
        cin >> u >> v;
        out[u] ++;
        in[v] ++;
        graph[u][v] = 1;
    }
    queue<int> handle;
    vector<int> ans;
    for (int i = 1 ; i <= n ; i ++) {
        if (in[i] == 0) {
            handle.push(i);
            ans.push_back(i);
        }
    }
    
    
    while (handle.size() > 0) {
        int node = handle.front();
        for (int j = 1 ; j <= n ; j ++) {
            if (graph[node][j] == 1) {
                in[j] --;
                if (in[j] == 0) {
                    handle.push(j);
                    ans.push_back(j);
                }
            }
        }
    }
    
    if (ans.size() == n) {
        for (int i = 0 ; i < ans.size() ; i ++ ) {
            cout << ans[i] << " ";
        }
        cout << endl;
    } else {
        cout << "-1" << endl;
    }
    
    return 0;
}