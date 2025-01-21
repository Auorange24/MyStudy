#include <bits/stdc++.h>
using namespace std;

const int N = 1001;

int graph[N][N];

int dx[8] = {1, 2, 2, 1, -1, -2, -2, -1};
int dy[8] = {2, 1, -1, -2, -2, -1, 1, 2};

typedef pair<pair<int, int>, int> NextNode;

int cal_g(int tx, int ty, int ex, int ey) {
    return abs(tx - ex) * abs(tx - ex) + abs(ey - ty) * abs(ey -ty);
}

int traversal(int cur_x, int cur_y, int end_x, int end_y, int step) {
    cout << "cur_x = " << cur_x << "cur_y = " << cur_y << endl;
    if (cur_x == end_x && cur_y == end_y) {
        return step;
    }
    vector<NextNode> nodes;
    for (int i = 0 ; i < 8 ; i ++) {
        int tx = cur_x + dx[i];
        int ty = cur_y + dy[i];
        if (tx >= 1 && tx <= 1000 && ty >= 1 && ty <= 1000) {
            int g = cal_g(tx, ty, end_x, end_y);
            NextNode node;
            node.first.first = tx;
            node.first.second = ty;
            node.second = g + step + 1;
            nodes.push_back(node);
        }
    }

    for (int i = 0 ; i < nodes.size() ; i ++) {
        cout << "x = " << nodes[i].first.first << " y = " << nodes[i].first.second << "distance = " << nodes[i].second << endl;
    }

    sort(nodes.begin(), nodes.end(), [](NextNode a, NextNode b) {
        return a.second < b.second;
    });

    return traversal(nodes[0].first.first, nodes[0].first.second, end_x, end_y, step + 1);
}

int solve(int sx, int sy, int ex, int ey) {
    return traversal(sx, sy, ex, ey, 0);
}

int main() {
    int t;
    cin >> t;
    while (t --) {
        int sx, sy, ex, ey;
        cin >> sx >> sy >> ex >> ey;
        cout << solve(sx, sy, ex, ey) << endl;
    }
}