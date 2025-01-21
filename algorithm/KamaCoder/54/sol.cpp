#include <bits/stdc++.h>
using namespace std;

typedef pair<pair<int, int>, int> Node;
const int M = 1e5 + 10;
Node nodes[M];
const int N = 1e4 + 10;
int f[N];

int n, m;

void init() {
    for (int i = 0 ; i <= n ; i ++) {
        f[i] = i;
    }
}

int find(int x) {
    if (x == f[x]) {
        return x;
    } else {
        return find(f[x]);
    }
}

int main() {
    cin >> n >> m;
    init();
    for (int i = 0 ; i < m ; i ++) {
        int x, y, v;
        cin >> x >> y >> v;
        nodes[i].first.first = x;
        nodes[i].first.second = y;
        nodes[i].second = v;
    }
    sort(nodes, nodes + m, [](Node a, Node b) {
        return a.second < b.second;
    });
    
    int cur = 0;
    int ans = 0;
    
    for (int i = 0 ; i < m ; i ++) {
        if (find(nodes[i].first.first) == find(nodes[i].first.second)) {
            continue;
        } else {
            f[find(nodes[i].first.second)] = find(nodes[i].first.first);
            cur ++;
            ans += nodes[i].second;
            if (cur == n - 1) {
                break;
            }
        }
    } 
    cout << ans << endl;
    return 0;
}