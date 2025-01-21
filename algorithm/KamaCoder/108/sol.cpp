#include <bits/stdc++.h>

using namespace std;

const int N = 1010;
int f[N];
int n;

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
    cin >> n;
    int size = n;
    int ans_s, ans_e;
    init();
    // cout << "size = " << size << endl;
    for (int i = 0 ; i < n ; i ++) {
        int s, e;
        cin >> s >> e;
        // cout << "s = " << s << " e = " << e << endl;
        // cout << "find[s] = " << find(s) << "find[e] = " << find(e) << endl;
        if (find(s) == find(e)) {
            ans_s = s;
            ans_e = e;
        } else {
            f[find(s)] = find(e);
            size --;
        }
    }   
    cout << ans_s << " " << ans_e << endl;
    return 0;
}